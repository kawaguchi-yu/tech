FROM node
WORKDIR /usr/app
COPY . .
RUN apt-get update -y && apt-get upgrade -y && yarn install
RUN yarn build
EXPOSE 80:80
CMD ["yarn","start"]