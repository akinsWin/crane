package catalog

import (
	"errors"

	log "github.com/Sirupsen/logrus"
)

func (catalogApi *CatalogApi) Save(catalog *Catalog) error {
	count := 0
	if err := catalogApi.
		DbClient.
		Model(&Catalog{}).
		Where("name = ?", catalog.Name).
		Count(&count).Error; err != nil {
		log.Errorf("get catalog error: %v", err)
		return err
	}

	if count > 0 {
		log.Warnf("catlog: %s already exist", catalog.Name)
		return errors.New("already exist")
	}

	if err := catalogApi.DbClient.Save(catalog).Error; err != nil {
		log.Errorf("save catalog error: %v", err)
		return err
	}

	return nil
}

func (catalogApi *CatalogApi) List() ([]Catalog, error) {
	var catalogs []Catalog
	err := catalogApi.DbClient.Find(&catalogs).Error
	return catalogs, err
}

func (catalogApi *CatalogApi) Get(catalogId uint64) (Catalog, error) {
	var catalog Catalog
	err := catalogApi.DbClient.Where("id = ?", catalogId).First(&catalog).Error
	return catalog, err
}

func (catalogApi *CatalogApi) Delete(catalogId uint64) error {
	return catalogApi.DbClient.Delete(&Catalog{ID: catalogId}).Error
}

func (catalogApi *CatalogApi) Update(catalog *Catalog) error {
	count := 0
	if err := catalogApi.DbClient.
		Model(&Catalog{}).
		Where("name = ? AND id != ? AND user_id = ?", catalog.Name, catalog.ID, catalog.UserId).
		Count(&count).
		Error; err != nil {
		log.Errorf("get catalog error: %v", err)
		return err
	}

	if count > 0 {
		log.Warnf("catlog: %s already exist", catalog.Name)
		return errors.New("already exist")
	}
	return catalogApi.DbClient.Save(catalog).Error
}
