set -e

cd frontend
yarn build
zip -r build.zip build
scp build.zip $DEPLOY_USERNAME@$DEPLOY_HOSTNAME:~

ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "rm -rf /var/www/html/*"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT 'unzip -o ~/build.zip'
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "mv ~/build/* /var/www/html/"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT "mv ~/build/.htaccess /var/www/html/"
ssh $DEPLOY_USERNAME@$DEPLOY_HOSTNAME -p $DEPLOY_SSH_PORT 'rm ~/build.zip'
