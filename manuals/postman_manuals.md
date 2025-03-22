
# Postmanの設定

まずは `GetAllUsers` の設定を行う。

- 1. New Collection
  - gRPC basic用で追加すること
- 2. 以下設定を行う。
  - Service definition
    - `user.proto` を設定
  - address : localhost:50051
  - pb.UserService/GetAllUsers
  - Message
    - {} 
- 3. 認証JWTトークンを設定する。
  - Authorization
    - `Bearer Token` に設定し、`Token`を設定すること。 
- 4. Invokeボタンを押下し、バックエンドWebAPIへリクエストを投げる。

## GetAllTodos

- message

```json
{}
```

## GetTodoById

- message

```json
{
    "id": ""
}
```

## GetTodoByUserId

- message

```json
{
    "userId": ""
}
```

## Createtodo

- message

```json
{
    "description": "",
    "userId": ""
}
```

## Updatetodo

- message

```json
{
    "id": "",
    "description": "",
    "completed": true,
    "userId": ""
}
```

## DeleteTodo

- message

```json
{
    "id": ""
}
```

## Login

- `Header`から`Authorization`を外すこと。

- message

```json
{
    "email": "",
    "passwordßß": ""
}
```

