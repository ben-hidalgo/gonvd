FROM node:10.15.3-jessie

WORKDIR /tests

ADD package.json package.json
ADD package-lock.json package-lock.json

RUN npm i

COPY . .

CMD ["npm", "test"]
