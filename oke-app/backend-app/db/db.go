package db

import (
	"fmt"
	"os"
	"time"

	"github.com/tniita/ochacafe-github-actions/oke-app/backend-app/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbInfo() string {
	dbInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	return dbInfo
}

func SetupDB() {
	dbInfo := GetDbInfo()
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database" + err.Error())
	}

	db.AutoMigrate(&repo.Items{})

	Items := []*repo.Items{
		{
			Name:       "Kubernetes超入門",
			Date:       time.Date(2023, 6, 7, 19, 00, 00, 000000, time.UTC).Format("20060102150405"),
			Topics:     "Kubernetes",
			Presenters: "Yutaka Ichikawa",
		},
		{
			Name:       "IaCのベストプラクティス",
			Date:       time.Date(2023, 7, 5, 19, 00, 00, 000000, time.UTC).Format("20060102150405"),
			Topics:     "Terraform, Pulumi",
			Presenters: "Shuhei Kawamura",
		},
		{
			Name:       "GitHub Actionsを使いこなせ！",
			Date:       time.Date(2023, 8, 9, 19, 00, 00, 000000, time.UTC).Format("20060102150405"),
			Topics:     "GitHub Actions",
			Presenters: "Takuya Niita",
		},
		{
			Name:       "セキュアなWeb APIの作り方",
			Date:       "TBD",
			Topics:     "Keycloak, cert-manager",
			Presenters: "Shuhei Kawamura",
		},
		{
			Name:       "Cluster API - K8sクラスタ管理の新スタイル",
			Date:       "TBD",
			Topics:     "Cluster API, API Provider for OCI",
			Presenters: "Takuya Niita",
		},
		{
			Name:       "次はこれでしょ！eBPF",
			Date:       "TBD",
			Topics:     "eBPF, CNI, Cilium",
			Presenters: "Yutaka Ichikawa",
		},
	}

	db.Create(Items)
}
