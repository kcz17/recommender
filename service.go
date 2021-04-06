package recommender

// service.go contains the definition and implementation (business logic) of the
// recommender service. Everything here is agnostic to the transport (HTTP).

import (
	"errors"
	randexp "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"

	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
)

// Service is the recommender service, providing read operations on a saleable
// recommender of sock products.
type Service interface {
	Get() (Sock, error) // GET /recommender
	Health() []Health   // GET /health
}

// Middleware decorates a Service.
type Middleware func(Service) Service

// Sock describes the thing on offer in the catalogue.
type Sock struct {
	ID          string   `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	ImageURL    []string `json:"imageUrl" db:"-"`
	ImageURL_1  string   `json:"-" db:"image_url_1"`
	ImageURL_2  string   `json:"-" db:"image_url_2"`
	Price       float32  `json:"price" db:"price"`
	Count       int      `json:"count" db:"count"`
	Tags        []string `json:"tag" db:"-"`
	TagString   string   `json:"-" db:"tag_name"`
}

// Health describes the health of a service
type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// ErrNotFound is returned when there is no sock for a given ID.
var ErrNotFound = errors.New("not found")

// ErrDBConnection is returned when connection with the database fails.
var ErrDBConnection = errors.New("database connection error")

// NewRecommenderService returns an implementation of the Service interface,
// with connection to an SQL database.
func NewRecommenderService(db *sqlx.DB, logger log.Logger) Service {
	return &newsService{
		db:     db,
		logger: logger,
	}
}

type newsService struct {
	db     *sqlx.DB
	logger log.Logger
}

func (s *newsService) Get() (Sock, error) {
	var socks []Sock
	query := "SELECT sock.sock_id AS id, sock.name, sock.description, sock.price, sock.count, sock.image_url_1, sock.image_url_2, GROUP_CONCAT(tag.name) AS tag_name FROM sock JOIN sock_tag ON sock.sock_id=sock_tag.sock_id JOIN tag ON sock_tag.tag_id=tag.tag_id GROUP BY id;"

	err := s.db.Select(&socks, query)
	if err != nil {
		s.logger.Log("database error", err)
		return Sock{}, err
	}
	for i, s := range socks {
		socks[i].ImageURL = []string{s.ImageURL_1, s.ImageURL_2}
		socks[i].Tags = strings.Split(s.TagString, ",")
	}

	// Set the random seed to the current time for sufficient uniqueness.
	randSeed := uint64(time.Now().UTC().UnixNano())
	delay := distuv.Normal{
		Mu:    2,
		Sigma: 2,
		Src:   randexp.NewSource(randSeed),
	}.Rand()
	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

	if len(socks) > 0 {
		return socks[rand.Intn(len(socks))], nil
	}
	return Sock{}, nil
}

func (s *newsService) Health() []Health {
	var health []Health
	dbstatus := "OK"

	err := s.db.Ping()
	if err != nil {
		dbstatus = "err"
	}

	app := Health{"recommender", "OK", time.Now().String()}
	db := Health{"recommender-db", dbstatus, time.Now().String()}

	health = append(health, app)
	health = append(health, db)

	return health
}
