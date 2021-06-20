docker build -t ranggarifqi/qashir-backend .
# Install heroku CLI
curl https://cli-assets.heroku.com/install.sh | sh
# Login to heroku
heroku auth:token
# Deployment
docker login --username=_ --password=$(heroku auth:token) registry.heroku.com
docker tag ranggarifqi/qashir-backend registry.heroku.com/qashir-backend/web
docker push registry.heroku.com/qashir-backend/web
heroku container:release web --app=qashir-backend