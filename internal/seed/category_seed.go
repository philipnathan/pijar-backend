package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	repo "github.com/philipnathan/pijar-backend/internal/category/repository"
	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	categories := []model.Category{
		{
			Category_name: "Pertanian",
			Image_url:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQOD9cAXPUyh6yJu8ThKALMyB3ROQTmJv4_uR9J9NDL1GdnUZjfQuPuRXpMMM8CuF3t-Wg&usqp=CAU"},
		{
			Category_name: "Kewirausahaan",
			Image_url:     "https://st2.depositphotos.com/1010613/11110/i/450/depositphotos_111105200-stock-photo-businesspeople-making-various-business-chart.jpg"},
		{
			Category_name: "Kerajinan Tangan",
			Image_url:     "https://cdn0-production-images-kly.akamaized.net/AHPg4s_tV7B2FgD-6VsRrzniHEE=/500x0/smart/filters:quality(75):strip_icc():format(webp)/kly-media-production/medias/4342502/original/049461600_1677665012-Berburu_Kerajinan_Tangan_di_INACRAFT_2023-Angga-1.jpg"},
		{
			Category_name: "Pengolahan Hasil Pertanian",
			Image_url:     "https://3.bp.blogspot.com/-rp4EgZWQmWs/WOVnoPJey4I/AAAAAAAAADE/bnWhUXddIwgeS8fXY2SMkfSSGs3fVg5zQCLcB/s1600/sapta-usaha-tani-kesejahteraan-petani.jpg"},
		{
			Category_name: "Pemasaran Digital",
			Image_url:     "https://jatismobile.com/wp-content/uploads/2023/01/marketing-design-concept-with-laptop-loudspeaker-other-elements-isometric-vector-illustration-landing-page-template-web-background-1024x683.jpg"},
		{
			Category_name: "Perikanan & Peternakan",
			Image_url:     "https://asset.kompas.com/crops/S_oG_5SNKUKpAYMPPdLuwtEWd7k=/7x8:917x615/750x500/data/photo/2023/03/20/6418338060d55.jpg"},
		{
			Category_name: "Pemrograman Komputer",
			Image_url:     "https://kazokku.com/blog/wp-content/uploads/2023/07/contoh-bahasa-pemrograman.webp",
		},
		{
			Category_name: "Perdagangan Internasional",
			Image_url:     "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh2OYYe0d2g8APHSMdJVp5qQMtsxmTkQlcSSRcM01_4TNq4bznct_V9SAoGHQ8LM96184DxiVIIk_12dZMDLYXoAurMyEQ8ekwEg_ce0NsxcVNd-F8_i6-zPPAVbj9c5tPkmGlAitOnN14/s1600/Perdagangan+Internasional.jpg",
		},
		{
			Category_name: "Desain",
			Image_url:     "https://greatmind.id/uploads/article-detail/0589d0c62b4dd8ca5dab7035bed8c9b795c23bd4.jpeg",
		},
	}

	var count int64
	if err := db.Model(&model.Category{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, category := range categories {
		if err := repo.NewCategoryRepository(db).SaveCategory(&category); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
