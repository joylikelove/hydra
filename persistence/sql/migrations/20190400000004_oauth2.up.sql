ALTER TABLE hydra_oauth2_access ADD active BOOL NOT NULL DEFAULT TRUE;
ALTER TABLE hydra_oauth2_refresh ADD active BOOL NOT NULL DEFAULT TRUE;
ALTER TABLE hydra_oauth2_code ADD active BOOL NOT NULL DEFAULT TRUE;
ALTER TABLE hydra_oauth2_oidc ADD active BOOL NOT NULL DEFAULT TRUE;
ALTER TABLE hydra_oauth2_pkce ADD active BOOL NOT NULL DEFAULT TRUE;

