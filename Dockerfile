FROM node:14-alpine
WORKDIR /root
COPY . .
RUN ls
RUN yarn install
ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools wheel twine
RUN apk add openjdk11 git
RUN git config --global user.email "bot@cisco.com"
RUN git config --global user.name "piedPiperBot"
RUN npm install @openapitools/openapi-generator-cli -g
RUN openapi-generator-cli version
RUN openapi-generator-cli version-manager set 6.1.0
CMD ["node", "/root/build.sh"]
