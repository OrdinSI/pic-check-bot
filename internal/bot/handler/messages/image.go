package messages

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *MessageHandler) imageHandle(ctx context.Context, b *bot.Bot, update *models.Update) {
	photo := update.Message.Photo[len(update.Message.Photo)-1]
	photoID := photo.FileID

	fileInfo, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: photoID,
	})
	if err != nil {
		h.log.Error("Error in GetFile in ImageHandle:", err)
		return
	}

	fileData, err := h.downloadFileToMemory(fileInfo.FilePath, h.cfg.Token)
	if err != nil {
		h.log.Error("Failed to download file: %v", err)
		return
	}
	req := model.ImageRequest{
		GroupID:   update.Message.Chat.ID,
		UserID:    update.Message.From.ID,
		FileID:    photoID,
		MessageID: update.Message.ID,
	}

	image, userName, err := h.ucase.CheckHashImage(ctx, fileData, req)
	if err != nil {
		h.log.Error("Error in CheckHashImage in ImageHandle:", err)
		return
	}

	if image != nil {
		msg := fmt.Sprintf(
			"⚠️Внимание⚠️\nОбнаружен боян.\nПервоисточник: @%s\nДата: %s\nСсылка: %s",
			userName,
			image.PostTime.Format("2006-01-02"),
			generateLink(image.GroupID, image.MessageID),
		)
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   msg,
		})
		if err != nil {
			h.log.Error("Error in SendMessage in ImageHandle:", err)
			return
		}
	}

}

func (h *MessageHandler) downloadFileToMemory(filePath, token string) ([]byte, error) {
	h.log.Info("Downloading file to memory:", filePath, "token", token)
	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}
	h.log.Info("File downloaded successfully: Size=%d bytes", len(fileData))
	return fileData, nil
}

func generateLink(groupID int64, messageID int) string {
	chatID := getChatIDForLink(groupID)

	return fmt.Sprintf("https://t.me/c/%s/%d", chatID, messageID)
}

func getChatIDForLink(chatID int64) string {
	if chatID < 0 {
		return strconv.FormatInt(-chatID, 10)[3:]
	}
	return strconv.FormatInt(chatID, 10)
}
