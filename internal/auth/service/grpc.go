package service

import (
    "context"
    "errors"
    "time"

    "github.com/sxd0/SweetTweet/internal/auth/model"
    "github.com/sxd0/SweetTweet/internal/auth/repository"
    pb "github.com/sxd0/SweetTweet/proto/authpb"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
    pb.UnimplementedAuthServiceServer
    repo *repository.UserRepository
    jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
    return &AuthService{
        repo: repo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Email: req.Email,
        PasswordHash: string(hash),
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }

    return &pb.RegisterResponse{Message: "user registered"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    user, err := s.repo.GetByEmail(req.Email)
    if err != nil {
        return nil, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": user.Email,
        "user_id": user.ID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return nil, err
    }

    return &pb.LoginResponse{Token: tokenString}, nil
}

func (s *AuthService) Me(ctx context.Context, req *pb.MeRequest) (*pb.MeResponse, error) {
    token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.jwtSecret), nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid or expired token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("cannot parse token claims")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return nil, errors.New("user_id not found in token")
    }

    email, ok := claims["email"].(string)
    if !ok {
        return nil, errors.New("email not found in token")
    }

    return &pb.MeResponse{
        UserId: int64(userIDFloat),
        Email:  email,
    }, nil
}

