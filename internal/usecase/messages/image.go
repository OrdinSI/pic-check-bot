package messages

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/OrdinSI/pic-check-bot/internal/usecase"
	"github.com/corona10/goimagehash"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/vp8"
	_ "golang.org/x/image/vp8l"
	_ "golang.org/x/image/webp"
)

func (m *messages) CheckHashImage(ctx context.Context, fileData []byte, req model.ImageRequest) (*model.Image, string, error) {
	m.log.Info("Received file data size: %d bytes", len(fileData))
	reader := bytes.NewReader(fileData)
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}
	m.log.Info("Image format: %s", format)
	hash, err := goimagehash.PerceptionHash(img)
	if err != nil {
		return nil, "", fmt.Errorf("failed to calculate hash: %w", err)
	}

	fileHash, err := serializeHash(hash)
	if err != nil {
		return nil, "", fmt.Errorf("failed to serialize hash: %w", err)
	}

	hashPart1 := binary.LittleEndian.Uint64(fileHash[:8])
	hashPart2 := binary.LittleEndian.Uint64(fileHash[8:16])

	var candidates []model.Image
	if err := m.repoM.GetImagesByHashParts(ctx, &candidates, hashPart1, hashPart2); err != nil {
		return nil, "", fmt.Errorf("failed to get images by hash part: %w", err)
	}

	for _, candidate := range candidates {
		hashObj, err := deserializeHash(candidate.FileHash)
		if err != nil {
			m.log.Error("Failed to deserialize hash:", err)
			continue
		}

		distance, err := hashObj.Distance(hash)
		if err != nil {
			m.log.Error("Failed to calculate distance:", err)
			continue
		}
		if distance < usecase.DuplicateThreshold {
			if err := m.processedDuplicateImage(ctx, req, candidate.ID); err != nil {
				m.log.Error("Failed to process duplicate image:", err)
				continue
			}
			user, err := m.repoU.GetUser(ctx, candidate.UserID)
			if err != nil {
				m.log.Error("Failed to get user:", err)
				continue
			}
			return &candidate, user.Username, nil
		}
	}

	newImages := &model.Image{
		HashPart1: hashPart1,
		HashPart2: hashPart2,
		FileHash:  fileHash,
		UserID:    req.UserID,
		GroupID:   req.GroupID,
		FileID:    req.FileID,
		MessageID: req.MessageID,
	}

	if err := m.repoM.CreateImage(ctx, newImages); err != nil {
		return nil, "", err
	}

	return nil, "", nil
}

func (m *messages) processedDuplicateImage(ctx context.Context, req model.ImageRequest, imageID uint) error {
	if err := m.repoM.CreateRepost(ctx, &model.Repost{
		ImageID: imageID,
		UserID:  req.UserID,
		GroupID: req.GroupID,
	}); err != nil {
		return err
	}
	return nil
}

func serializeHash(hash *goimagehash.ImageHash) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	err := hash.Dump(writer)
	if err != nil {
		return nil, err
	}
	writer.Flush()
	return buffer.Bytes(), nil
}

func deserializeHash(data []byte) (*goimagehash.ImageHash, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	return goimagehash.LoadImageHash(reader)
}
