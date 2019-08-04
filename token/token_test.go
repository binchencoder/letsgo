package token

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJwtToken(t *testing.T) {
	signingKey := "TestSigningKey_1fg4567yu"

	info := TokenClientInfo{
		ClientId: "TestClientId",
		CorpCode: "TestCid",
		Scopes:   []string{"scope1", "scope2"},
	}

	Convey("Create and validate a correct signed token", t, func() {
		tokenStr, err := CreateSignedJwtToken(signingKey, &info, 1000)
		So(err, ShouldBeNil)
		So(tokenStr, ShouldNotBeNil)

		token, err := ValidateSignedJwtToken(tokenStr, signingKey)
		So(err, ShouldBeNil)
		So(token, ShouldNotBeNil)

		So(info.ClientId, ShouldEqual, token.ClientId)
		So(info.CorpCode, ShouldEqual, token.CorpCode)
		So(reflect.DeepEqual(info.Scopes, token.Scopes), ShouldBeTrue)
	})

	Convey("Create and validate an expired signed token", t, func() {
		tokenStr, err := CreateSignedJwtToken(signingKey, &info, -1000)
		So(err, ShouldBeNil)
		So(tokenStr, ShouldNotBeNil)

		token, err := ValidateSignedJwtToken(tokenStr, signingKey)
		So(err, ShouldNotBeNil)
		So(token, ShouldBeNil)
	})
}
