package services

import (
	"context"
	"job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (s *Service) StoreCompany(ctx context.Context, newCompany models.NewCompany) (models.Company, error) {

	cDetails := models.Company{
		CompanyName:     newCompany.CompanyName,
		Location: newCompany.Location,
	}
	companyData, err := s.UserRepo.CreateCompany(ctx, cDetails)
	if err != nil {
		log.Info().Err(err).Send()
		return models.Company{}, err
	}
	return companyData, nil
}

func (s *Service) GetCompanyData(ctx context.Context, companyId uint64) (models.Company, error) {
	companyData, err := s.UserRepo.ViewCompanyById(ctx, companyId)
	if err != nil {
		log.Info().Err(err).Send()
		return models.Company{}, err
	}
	return companyData, nil
}

func (s *Service) GetAllCompanyData(ctx context.Context) ([]models.Company, error) {
	companyDetails, err := s.UserRepo.ViewCompanies(ctx)
	if err != nil {
		log.Info().Err(err).Send()
		return nil, err
	}
	return companyDetails, nil
}
