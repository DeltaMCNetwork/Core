package server

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://sessionserver.mojang.com/session/minecraft/hasJoined?username="

const (
	AuthSuccess Result = iota
	AuthFail
	AuthError
)

type Result byte

type AuthenticationResult struct {
	Result     Result
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Properties *[]Properties `json:"properties"`
}

type Properties struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

type IAuthenticator interface {
	Authenticate(IPlayer, *MinecraftServer, []byte) *AuthenticationResult
}

type MojangAuthenticator struct {
	httpClient *http.Client
}

func CreateMojangAuthenticator() *MojangAuthenticator {
	return &MojangAuthenticator{httpClient: http.DefaultClient}
}

func (ma *MojangAuthenticator) Authenticate(player IPlayer, server *MinecraftServer, secret []byte) *AuthenticationResult {
	// Authenticate player with Mojang servers

	resp, err := ma.httpClient.Get(url + player.GetUsername() + "&serverId=" + makeHash(secret, server) + "&ip=" + player.GetIP())

	authResult := &AuthenticationResult{}

	if err != nil {
		Error("Error authenticating player with Mojang session server: " + err.Error())
		authResult.Result = AuthError
		return authResult
	}

	if resp.StatusCode != http.StatusOK {
		authResult.Result = AuthFail
		return authResult
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		Error("Error reading response from Mojang session server: " + err.Error())
		authResult.Result = AuthError
		return authResult
	}

	err = json.Unmarshal(data, authResult)
	if err != nil {
		Error("Error unmarshalling JSON response from Mojang session server: " + err.Error())
		authResult.Result = AuthError
		return authResult
	}

	fmt.Println(authResult)

	return authResult
}

var _ IAuthenticator = (*MojangAuthenticator)(nil)

func makeHash(secret []byte, server *MinecraftServer) string {
	sha := sha1.New()

	sha.Write(secret)
	sha.Write(server.GetKeypair().Public.Key)
	hash := sha.Sum(nil)

	negative := (hash[0] & 0x80) == 0x80
	if negative {
		hash = twosComplement(hash)
	}

	// Trim away zeroes
	res := strings.TrimLeft(hex.EncodeToString(hash), "0")
	if negative {
		res = "-" + res
	}

	return res
}

func twosComplement(p []byte) []byte {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = ^p[i]
		if carry {
			carry = p[i] == 0xff
			p[i]++
		}
	}
	return p
}
