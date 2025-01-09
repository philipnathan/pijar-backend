package seed

import (
	"log"
	"time"

	model "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
)

func SeedSession(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	sessions := []model.MentorSession{
		{
			UserID:           1,
			CategoryID:       1,
			Title:            "Teknik Bertani Modern",
			ShortDescription: "Belajar teknik bertani modern untuk meningkatkan hasil panen.",
			Detail:           "Sesi ini akan membahas penggunaan pupuk organik, irigasi modern, dan pemanfaatan teknologi dalam pertanian.",
			Schedule:         time.Date(2024, 12, 22, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/6Zt7HZbRbe8/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLBaNkQoCRq5CgByzkb5KY_DzGUKrg",
			Link:             "https://example.com/session_1",
		},
		{
			UserID:           1,
			CategoryID:       2,
			Title:            "Cara Memulai Usaha Kecil",
			ShortDescription: "Langkah-langkah memulai usaha kecil di desa",
			Detail:           "Sesi ini akan membahas cara membuat rencana bisnis, strategi pemasaran, dan pengelolaan keuangan sederhana.",
			Schedule:         time.Date(2024, 12, 25, 14, 0, 0, 0, time.UTC),
			EstimateDuration: 90,
			ImageURL:         "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgPypTe2QvUkLKqazJHgVm5LCGeGgXwuUEHo3Z-pCK_a2e6xF1KKWAV1xwKnQ6FiPWvsNTe5YDXaENpba3ZaZu1PXEZFxnzOcM0lxgJuJWJa1q-GR-67QmjpADvU-t3Jaz1sgxaUTloYNiu/s1600/10+Tips+Memulai+Usaha+Kecil+yang+Bahkan+Belum+Anda+Dengar-min.png",
			Link:             "https://example.com/session_2",
		},
		{
			UserID:           3,
			CategoryID:       3,
			Title:            "Membuat Tas dari Bahan Daur Ulang",
			ShortDescription: "Pelatihan membuat tas dari plastik bekas dan kain perca.",
			Detail:           "Sesi ini akan membahas teknik daur ulang plastik bekas menjadi produk kreatif, cara menjahit tas, dan pemasaran hasil kerajinan.",
			Schedule:         time.Date(2024, 12, 28, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 150,
			ImageURL:         "https://d39wptbp5at4nd.cloudfront.net/media/96524_original_post-88993-32b095fb-e329-4cf3-a04e-0629ac98e5ca-2020-07-13t17-59-31.943-07-00.jpg",
			Link:             "https://example.com/session_3",
		},
		{
			UserID:           3,
			CategoryID:       4,
			Title:            "Membuat Keripik Pisang Berkualitas Ekspor",
			ShortDescription: "Cara membuat keripik pisang yang renyah dan tahan lama.",
			Detail:           "Pelatihan ini mencakup pemilihan bahan baku, teknik penggorengan, penggunaan alat vakum, dan pengemasan untuk pasar ekspor.",
			Schedule:         time.Date(2024, 12, 30, 15, 0, 0, 0, time.UTC),
			EstimateDuration: 180,
			ImageURL:         "https://static.promediateknologi.id/crop/0x0:0x0/750x500/webp/photo/p1/995/2024/12/13/WhatsApp-Image-2024-12-12-at-221459_90f678e5-1808357494.jpg",
			Link:             "https://example.com/session_4",
		},
		{
			UserID:           1,
			CategoryID:       1,
			Title:            "Teknik Bertani Rumahan",
			ShortDescription: "Belajar teknik bertani rumahan untuk meningkatkan produksi tanaman.",
			Detail:           "Sesi ini akan membahas penggunaan pupuk organik, perawatan tanaman, dan pemanfaatan teknologi dalam pertanian.",
			Schedule:         time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/6Zt7HZbRbe8/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLBaNkQoCRq5CgByzkb5KY_DzGUKrg",
			Link:             "https://example.com/session_1",
		},
		{
			UserID:           3,
			CategoryID:       1,
			Title:            "Teknik Bertani di Indonesia",
			ShortDescription: "Belajar teknik bertani di Indonesia untuk meningkatkan produksi tanaman.",
			Detail:           "Sesi ini akan membahas penggunaan pupuk organik, perawatan tanaman, dan pemanfaatan teknologi dalam pertanian.",
			Schedule:         time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/6Zt7HZbRbe8/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLBaNkQoCRq5CgByzkb5KY_DzGUKrg",
			Link:             "https://example.com/session_1",
		},
	}

	var count int64
	if err := db.Model(&model.MentorSession{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, session := range sessions {
		if err := tx.Create(&session).Error; err != nil {
			tx.Rollback()
			log.Printf("Error seeding session: %v", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil
}
