package auth

import (
	"ddd/common/logging"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateToken結果用の構造体
type GenerateTokenResults struct {
	Token string
	Jti   string
}

// ParseToken結果用の構造体
type ParseTokenExtractClaimsAnalysis struct { // クレーム部分
	Token *jwt.Token
	Id    string
	Jti   string
	Exp   time.Time
}
type Errs struct {
	InputErr    error
	InternalErr error
}

// ユーザーidで認証トークンを生成
func GenerateToken(userUuid string) (GenerateTokenResults, error) {
	// 返り血用の構造体セット
	var results GenerateTokenResults
	var err error

	// uuidを作成
	newJti, err := uuid.NewRandom() // 新しいuuidの生成
	if err != nil {
		return GenerateTokenResults{}, err
	}
	results.Jti = newJti.String()

	// // // テーブルを更新。
	// // if err := model.SaveJti(userUuid, newJti.String()); err != nil { // Userテーブルを更新
	// // 	return "", "", err
	// // }

	// クレーム部分 // ParseToken結果用の構造体
	claims := jwt.MapClaims{
		"id":  userUuid,    // user_uuid  // クレーム内は1単語のみとしている
		"jti": results.Jti, // new jti_uuid
		"exp": time.Now().Add(time.Second * time.Duration(jwtTokenLifeTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)    // トークン(JWTを表す構造体)作成
	results.Token, err = token.SignedString([]byte(jwtSecretKey)) // []byte()でバイト型のスライスに変換し、SignedStringで署名。
	if err != nil {
		return GenerateTokenResults{}, err
	}

	return results, nil
}

// トークン解析検証
// 成功時に得られる分析結果と複数のエラーが返るので、それぞれ構造体として扱い構造体で返す
// エラーは入力値が不正な場合と処理エラーなどが考えられ、それぞれフィールドとして定義し、呼び出し側では構造体のフィールドを!=nilでチェックしハンドルする

// トークンから必要情報取得
func ParseTokenExtractClaims(tokenString string) (ParseTokenExtractClaimsAnalysis, Errs) {
	// 返り血用の構造体セット
	var analysis ParseTokenExtractClaimsAnalysis
	var errs Errs
	// あらかじめanalysisを宣言したため、analysis, err := ができないことへの対策
	var err error

	// 署名が正しければ、解析用の鍵を使う。(無名関数内で署名方法がHMACであるか確認し、HMACであれば秘密鍵を渡し、jwtトークンを解析する。)
	analysis.Token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 署名を確認
			logging.ErrorLog(fmt.Sprintf("Unexpected signature method: %v.", token.Header["alg"]), nil)
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil // 署名が正しければJWT_SECRET_KEYをバイト配列にして返す
	})
	if err != nil {
		errs.InternalErr = err
		return analysis, errs
	}

	// 下のクレーム検証処理(:elseスコープ内)で持ち出したい値をあらかじめ宣言しておく。
	//var id string // id
	// 構造体で管理することで不要に！

	// トークン自体が有効か秘密鍵を用いて確認。また、クレーム部分も取得。(トークンの署名が正しいか、有効期限内か、ブラックリストでないか。)
	claims, ok := analysis.Token.Claims.(jwt.MapClaims) // MapClaimsにアサーション
	if !ok || !analysis.Token.Valid {                   // 取得に失敗または検証が失敗
		errs.InputErr = errors.New("invalid authentication token")
		return analysis, errs
	}

	// 結果用変数に入れる 検証は呼び出しもとで

	// id
	analysis.Id, ok = claims["id"].(string) // goではJSONの数値は少数もカバーしたfloatで解釈される
	if !ok {
		errs.InputErr = errors.New("id could not be obtained from the token")
		return analysis, errs
	}
	// jti
	analysis.Jti, ok = claims["jti"].(string)
	if !ok {
		errs.InputErr = errors.New("jti could not be obtained from the token")
		return analysis, errs
	}
	// expを検証
	expAsserted, ok := claims["exp"].(float64)
	if !ok {
		errs.InputErr = errors.New("exp could not be obtained from the token")
		return analysis, errs
	}
	analysis.Exp = time.Unix(int64(expAsserted), 0) // Unix 時刻を日時に変換

	// 正常に終われば解析されたトークンとidを渡す。
	return analysis, errs
}
