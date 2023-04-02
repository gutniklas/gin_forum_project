# 基础镜像
FROM node:latest as build-stage

# 创建工作目录
WORKDIR /app

# 复制代码
COPY . .
# 安装依赖
RUN npm install

RUN npm audit fix

# 编译 Vue 应用
RUN npm run build

# 构建 Nginx 镜像
FROM nginx:latest
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
