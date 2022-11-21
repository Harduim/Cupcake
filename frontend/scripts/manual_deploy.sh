set -e

yarn build
zip -r build.zip build
scp -P $DEPLOY_SSH_PORT build.zip $DEPLOY_USERNAME@$DEPLOY_HOSTNAME:~

ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "rm -rf /var/www/html/*"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT 'unzip -o ~/build.zip'
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "mv ~/build/* /var/www/html/"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "mv ~/build/.htaccess /var/www/html/"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT 'rm ~/build.zip'

scp -P $DEPLOY_SSH_PORT public/flags/*.svg $DEPLOY_USERNAME@$DEPLOY_HOSTNAME:/var/www/html/flags/
