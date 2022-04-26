@REM docker build -t admin_api:v1 -f admin/cmd/api/Dockerfile .
@REM docker build -t classroom_api:v1 -f classroom/cmd/api/Dockerfile .
@REM docker build -t classroom_rpc:v1 -f classroom/cmd/rpc/Dockerfile .
@REM docker build -t excel_rpc:v1 -f excel/cmd/rpc/Dockerfile .
@REM docker build -t route_api:v1 -f route/cmd/api/Dockerfile .
docker build -t schedule_api:v1 -f schedule/cmd/api/Dockerfile .
@REM docker build -t student_rpc:v1 -f student/cmd/rpc/Dockerfile .
docker build -t student_api:v1 -f student/cmd/api/Dockerfile .
@REM docker build -t teacher_api:v1 -f teacher/cmd/api/Dockerfile .
@REM docker build -t teacher_rpc:v1 -f teacher/cmd/rpc/Dockerfile .
@REM docker build -t user_rpc:v1 -f user/cmd/rpc/Dockerfile .
@REM docker build -t user_api:v1 -f user/cmd/api/Dockerfile .