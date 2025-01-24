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
			Title:            "Teknik Pertanian Kopi Modern",
			ShortDescription: "Belajar teknik bertani kopi dari Brazil",
			Detail:           "Sesi ini akan membahas keseluruhan proses terkait dengan teknik pertanian kopi dari negara Brazil.",
			Schedule:         time.Date(2025, 02, 05, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/6Zt7HZbRbe8/hq720.jpg?sqp=-oaymwEhCK4FEIIDSFryq4qpAxMIARUAAAAAGAElAADIQj0AgKJD&rs=AOn4CLBaNkQoCRq5CgByzkb5KY_DzGUKrg",
			Link:             "https://www.youtube.com/embed/OwoKxFL43Ec?si=VHatLgKwSu2kujde",
		},
		{
			UserID:           1,
			CategoryID:       2,
			Title:            "Cara Memulai Bisnis dari 0",
			ShortDescription: "Langkah-langkah memulai usaha tanpa modal sama sekali.",
			Detail:           "Sesi ini akan bagaimana kita bisa memulai usaha tanpa modal sama sekali.",
			Schedule:         time.Date(2024, 11, 01, 14, 0, 0, 0, time.UTC),
			EstimateDuration: 90,
			ImageURL:         "https://i.ytimg.com/vi/8g1H3ipDNOs/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLAXJ2f5n8tkSl5i7y89fQYvUOfQ0A",
			Link:             "https://www.youtube.com/embed/8g1H3ipDNOs?si=Z2Dq37yS2RYdstd6",
		},
		{
			UserID:           3,
			CategoryID:       3,
			Title:            "Membuat Tas dari Bahan Daur Ulang",
			ShortDescription: "Pelatihan membuat tas dari plastik bekas bungkus kopi",
			Detail:           "Sesi ini akan membahas teknik daur ulang plastik bekas menjadi produk kreatif, dan cara menyatukannya menjadi tas yang dapat digunakan",
			Schedule:         time.Date(2025, 02, 05, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 150,
			ImageURL:         "https://i.ytimg.com/vi/bDus7YXtOIg/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLCPapIbHIJQi5f6ejRrsa8l3rq9kw",
			Link:             "https://www.youtube.com/embed/bDus7YXtOIg?si=CMsnXDrQuGRsZS5T",
		},
		{
			UserID:           3,
			CategoryID:       4,
			Title:            "Membuat Keripik Pisang",
			ShortDescription: "Cara membuat keripik pisang yang renyah dan tahan lama.",
			Detail:           "Pelatihan ini mencakup pemilihan bahan baku, dan teknik penggorengan agar keripik pisang yang renyah dan tahan lama.",
			Schedule:         time.Date(2025, 01, 30, 15, 0, 0, 0, time.UTC),
			EstimateDuration: 180,
			ImageURL:         "https://i.ytimg.com/vi/JNthbCOx3BI/hqdefault.jpg?sqp=-oaymwEnCOADEI4CSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLDZlyQif4Ch6gYQzCR6lIgnhVVarg",
			Link:             "https://www.youtube.com/embed/JNthbCOx3BI?si=gk977VRoH6zT5bKK",
		},
		{
			UserID:           1,
			CategoryID:       1,
			Title:            "Cara Menanam Kangkung Hidroponik di Ember",
			ShortDescription: "Belajar bagaimana cara menanam kangkung hidroponik di ember dengan mudah dan biaya yang murah.",
			Detail:           "Sesi ini akan membahas bahan-bahan yang diperlukan dan bagaimana cara menanam kangkung hidroponik di ember yang tentunya mudah dan murah.",
			Schedule:         time.Date(2024, 11, 1, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/-GyJN1tr9RM/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLDTi2xDla1Du_8kmpQnVwStmsiclA",
			Link:             "https://www.youtube.com/embed/-GyJN1tr9RM?si=77INhHqTXvuIvhkc",
		},
		{
			UserID:           3,
			CategoryID:       1,
			Title:            "Cara Menaikan Harga Jual Produk Pertanian",
			ShortDescription: "4 Cara menaikkan harga jual produk hasil pertanian supaya lebih tinggi dari sebelumnya.",
			Detail:           "Belajar berbagai cara untuk meningkatkan harga jual pertanian di masyarakat. Materi ini meruapakan hasil pengalaman penyuluh dalam berwirausaha bidang pertanian.",
			Schedule:         time.Date(2024, 12, 1, 9, 0, 0, 0, time.UTC),
			EstimateDuration: 120,
			ImageURL:         "https://i.ytimg.com/vi/4TvMrjzVQBY/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLA1M2pFTEsxnKk5mn25a-z-bkvWuA",
			Link:             "https://www.youtube.com/embed/4TvMrjzVQBY?si=VCe8ogkHOD2S3Ij7",
		},
		{
			UserID:           9,
			CategoryID:       7,
			Title:            "Mengintip Fitur-Fitur Baru CSS! | CSS Wrapped Part 1",
			ShortDescription: "Fitur-fitur baru pada CSS di 2024.",
			Detail:           "",
			Schedule:         time.Date(2025, 01, 28, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 150,
			ImageURL:         "https://media.licdn.com/dms/image/v2/D5612AQEJ_OtsZe59Cw/article-cover_image-shrink_720_1280/article-cover_image-shrink_720_1280/0/1731693543814?e=2147483647&v=beta&t=TIddd4WBiQHkZKtjLPTgUoXei_Y9wsBo2gBoLpOiCvA",
			Link:             "https://www.youtube.com/embed/D7HrMJBaPHw?si=0Ng2z6vo6x8jAY53",
		},
		{
			UserID:           9,
			CategoryID:       7,
			Title:            "Membuat Website Portfolio Menggunakan TAILWIND CSS 3 | NGOBAR #32",
			ShortDescription: "Membuat Website Portfolio Menggunakan TAILWIND CSS 3.",
			Detail:           "Di video ini, kita akan membuat website portfolio menggunakan TAILWIND CSS 3. Beberapa tools yang digunakan dalam membuat website ini adalah Tailwind CSS, HTML, dan CSS.",
			Schedule:         time.Date(2024, 12, 01, 15, 0, 0, 0, time.UTC),
			EstimateDuration: 180,
			ImageURL:         "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSOonb5anp0_jDqNtHBQYcqmEtYwfRticYTGg&s",
			Link:             "https://www.youtube.com/embed/8Ea4oq8qFtM?si=h4bvsQqEbVmwPrqE",
		},
		{
			UserID:           10,
			CategoryID:       8,
			Title:            "Apa itu Ekspor Impor dalam Perdagangan Internasional?",
			ShortDescription: "Penjelasan tentang ekspor impor dalam perdagangan internasional.",
			Detail:           "Materi yang akan dipelajari adalah bagaimana proses ekspor impor dalam perdagangan internasional. Materi ini akan diberikan secara ringkas dan mudah dipahami.",
			Schedule:         time.Date(2024, 12, 20, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 150,
			ImageURL:         "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRkEMh38pQzXLKHHLA2K_zugJqJlSjC1XBoFKMjQwNXXD9HMyZw",
			Link:             "https://www.youtube.com/embed/8kcxjcUgWyo?si=Fis_x0Q501lVDiUE",
		},
		{
			UserID:           10,
			CategoryID:       8,
			Title:            "Ekspor Impor Barang Jasa",
			ShortDescription: "Materi tentang ekspor impor barang jasa.",
			Detail:           "Materi yang akan dipelajari tentang ekspor impor barang jasa. Materi ini akan diberikan secara ringkas dan mudah dipahami.",
			Schedule:         time.Date(2025, 01, 30, 15, 0, 0, 0, time.UTC),
			EstimateDuration: 180,
			ImageURL:         "https://eximdotm.com/wp-content/uploads/2024/08/Jasa-Ekspor-5-1024x728.webp",
			Link:             "https://www.youtube.com/embed/w3Bf7Pa33gY?si=HGp6pouxNBBECbOj",
		},
		{
			UserID:           11,
			CategoryID:       9,
			Title:            "Desain Logo Untuk Usaha",
			ShortDescription: "Materi tentang desain logo untuk usaha.",
			Detail:           "Materi yang akan dipelajari tentang desain logo untuk usaha. Materi ini akan diberikan secara ringkas dan mudah dipahami.",
			Schedule:         time.Date(2025, 01, 28, 10, 0, 0, 0, time.UTC),
			EstimateDuration: 150,
			ImageURL:         "https://i.ytimg.com/vi/pYH5KiBLlhA/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLDiJDN2k5SSPaEyzHDk16p2OF27GQ",
			Link:             "https://www.youtube.com/embed/pYH5KiBLlhA?si=KTVHoMo-M4u99npm",
		},
		{
			UserID:           12,
			CategoryID:       5,
			Title:            "Pemasaran E-commerce di Tokopedia, Shopee, dan Tiktok",
			ShortDescription: "Materi tentang pemasaran e-commerce.",
			Detail:           "Materi yang akan dipelajari tentang pemasaran e-commerce. Materi ini akan diberikan secara ringkas dan mudah dipahami.",
			Schedule:         time.Date(2025, 01, 30, 15, 0, 0, 0, time.UTC),
			EstimateDuration: 180,
			ImageURL:         "https://i.ytimg.com/vi/0UCOFNNneas/hq720.jpg?sqp=-oaymwEnCNAFEJQDSFryq4qpAxkIARUAAIhCGAHYAQHiAQoIGBACGAY4AUAB&rs=AOn4CLBS3-Ad5vxJ-yrwNYhoEgFWgAzLFA",
			Link:             "https://www.youtube.com/embed/0UCOFNNneas?si=e5yCG-TzFoaRy5Ed",
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
