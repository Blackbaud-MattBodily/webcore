package services

import "encoding/xml"

//FTPCredentialsDTO is a data transfer object for retrieving a user's FTP
//Credentials.
type FTPCredentialsDTO struct {
	XMLName  xml.Name `xml:"FTPINFO"`
	UserName string   `json:"ftpUserName" xml:"FTPUSERNAME"`
	Password string   `json:"ftpPassword" xml:"FTPPASSWORD"`
}

//FTPService is a struct that contains an FTPRepository as well as functions
//for interacting with users' FTP credentials on the given FTPRepository.
type FTPService struct {
	repo FTPRepository
}

//FTPRepository is an interace that defines what functionality is required for
//an object to be considered a repository for the FTPService.
type FTPRepository interface {
	GetFTPCredentials(email string) (*FTPCredentialsDTO, error)
}

//NewFTPService returns a pointer to an FTPService instantiated with the given
//FTPRepository.
func NewFTPService(repo FTPRepository) *FTPService {
	return &FTPService{repo: repo}
}

//GetFTPCredentials retrieves a given user's (identified by email) FTP credentials.
func (f *FTPService) GetFTPCredentials(email string) (*FTPCredentialsDTO, error) {
	creds, err := f.repo.GetFTPCredentials(email)

	return creds, err
}
