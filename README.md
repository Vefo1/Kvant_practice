Go ������ API ��� ������������ ����������� ��������
��� Go API, ������� ������ ������ ��� �������� API ������������ ��������� ����������� �������� (HBA1C, LDLL, FERR, LDL, TG, HDL). ��� ��������� ������ � GET-���������� URL, ����������� �� � JSON-������ � ���������� POST-�������� �� ������� API, � ����� ���������� ����� �������� API �������. ������ ���������� � �������������� ������������������ �������, ����������� ���������������� ����, ����������� � ������������� ���������� ��� �����������.

������������
������-API ������� ����������� ����� ���������, ��������� � config.yml (app_auth.header_name, �� ��������� X-App-Auth), �� ���������, ��������� � config.yml (app_auth.token).

����� ��������� ��� ���� ������� (���� ���������):
uid (string, ����������� "web-client")

age (int)

gender (int)

1. HBA1C (/predict/hba1c)
���������: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/hba1c?age=30&gender=1&rdw=12.5&wbc=7.2&rbc=4.5&hgb=14.0&hct=42.0&mcv=90.0&mch=30.0&mchc=33.0&plt=250.0&neu=60.0&eos=2.0&bas=0.5&lym=30.0&mon=7.0&soe=10.0&chol=200.0&glu=90.0"




2. LDLL (/predict/ldll)
���������: chol, hdl, tg

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ldll?age=20&gender=1&chol=1&hdl=2&tg=3"




3. FERR (/predict/ferr)
���������: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, crp

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ferr?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&crp=7"




4. LDL (/predict/ldl)
���������: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/ldl?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8"




5. TG (/predict/tg)
���������: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/tg?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8"




6. HDL (/predict/hdl)
���������: rdw, wbc, rbc, hgb, hct, mcv, mch, mchc, plt, neu, eos, bas, lym, mon, soe, chol, glu

curl -X GET -H "X-App-Auth: m31A9Wo1=g#5c9z" "http://localhost:8080/predict/hdl?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=6&hct=5&mcv=4&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=77&glu=8"




��� �����������
������ API ����������� �� ����� ���������������� Go � �������������� ����������� ���������� net/http � ������� ����������� ������������� �� ��������� Go-��������.

���������� �������� ����������� � �� ����������:
config/:

�������� ����� � ���, ���������� �� �������� � ���������� ������������� ���������� (��������, �����, URL ������� API, ������ �����������, ������ �����������). ��� ��������� ����� ����������� ���������� ��� ��������� ��������� ����.

internal/models/:

���������� ��������� ������ (Go struct), ������� ������������ ��� ������������� ��������� � ����������, � ���������, ��� �������������� ������, ������������ �� ������� API, � �� ���������. ������������ ������� ��������� � �������� ������ � JSON.

internal/interfaces/:

�������� ����������� ����������� Go. ��� ���������� ��������� ��������� ��� ������-������ � ������ �����������, �������� �������� ���������� �� ���������� ����������. ��� ������������ �����������, ������������� � �������� �����������.

internal/services/:

��������� ������-������ ����������. ����� ��������� ������, ������� ��������� �������� ��������, ����� ��� �������������� � �������� API, ��������� ������ � ���������� ������-������. ������� ��������� ����������, ������������ � internal/interfaces/.

internal/handler/:

�������� HTTP-����������� (��������), ������� ��������� �������� HTTP-�������, ������ ���������, �������� ��������������� ������ �������� ��� ���������� ������-������ � ��������� HTTP-������ �������.

cmd/app/:

�������� ������ ����� � ����������. ���� ����� �������� �� ������������� ���� ����������� (������������, �������, ��������, ������������), ���������� ������������� ������������ (middleware) � ������ HTTP-�������.

����� ���������� �������:
������ ���������� GET ������ �� http://localhost:8080/predict/<model_name> (��������, /predict/ldll) � ����������� � URL � ���������� X-App-Auth.

������ �������� � main.go, ������� �������� ��� AuthMiddleware.

AuthMiddleware ��������� X-App-Auth �����. ���� �� ��������, ������ ����������� � 401 Unauthorized.

���� ����� ������������, AuthMiddleware �������� ������ ���������������� �������� (��������, LdllPredictHandler).

������� (����� handlePredictionRequest) ������ ��������� �� URL, ���������� �� � ��������� ��������������� ��������� models.XxxPredictRequest.

������� �������� ��������������� ����� (PredictXxx) � PredictServiceImpl.

����� PredictXxx (����� makePredictionRequest) ��������� POST-������ � �������� API (��������, https://apiml.labhub.online/api/v1/predict/ldll), ��������� ��������� (������� ����� ��� �������� API) � ���������� ���.

������� API ������������ ������ � ���������� �����.

PredictServiceImpl �������� ����� � ���������� ��� � �������.

������� �������� ����� (������-��� � ����) ������� �������.

��� ����� ���������� � ������� pkg/logger.