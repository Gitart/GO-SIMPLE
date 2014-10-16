func Test_Applications_GetAll(t *testing.T) {

    res := httptest.NewRecorder()
    req := makeReq("GET", "/v1/oauth2/applications")
    App().ServeHTTP(res, req)

    assert.Equal(t, res.Code, 200, "they should be equal")
}
