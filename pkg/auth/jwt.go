package auth

import (
  "fmt"
  "time"
  "math/rand"
  "errors"

  jwt "github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
  NewAccessToken(userId string, ttl time.Duration) (string, error)
  ParseToken(accessToken string) (string, error)
  NewRefreshToken() (string, error)
}

type JwtManager struct {
  signingKey string
} 

func NewJwtManager(signingKey string) (*JwtManager, error) {
  if signingKey == "" {
    return nil, errors.New("empty signing key")
  }
  
  return &JwtManager {
    signingKey: signingKey,
  }, nil
}

func (jm *JwtManager) NewAccessToken(userId string, ttl time.Duration) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims {
    ExpiresAt: time.Now().Add(ttl).Unix(),
    Subject: userId,
  })

  return token.SignedString(jm.signingKey) 
}

func (jm *JwtManager) ParseToken(accessToken string) (string, error) {
  token, err := jwt.Parse(accessToken, func (token *jwt.Token) (i any, err error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }    

    return []byte(jm.signingKey), nil
  })

  if err != nil {
    return "", err
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    return "", fmt.Errorf("error get user clims from token")
  }

  return claims["sub"].(string), nil
}

func (jm *JwtManager) NewRefreshToken() (string, error) {
  b := make([]byte, 32)

  s := rand.NewSource(time.Now().Unix())
  r := rand.New(s)

  if _, err := r.Read(b); err != nil {
    return "", err
  }

  return fmt.Sprintf("%x", b), nil
}


