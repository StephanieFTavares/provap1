import { Card, Input, Button, Typography } from "antd";
import { useState } from "react";
import api from "../services/api";

const { Title } = Typography;

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      await api.post("/login", { email, password });
      alert("Login realizado!");
    } catch {
      alert("Erro no login");
    }
  };

  return (
    <Card style={{ maxWidth: 400, margin: "20px auto" }}>
      <Title level={3}>Login</Title>

      <Input
        placeholder="Email"
        style={{ marginBottom: 10 }}
        onChange={(e) => setEmail(e.target.value)}
      />

      <Input.Password
        placeholder="Senha"
        style={{ marginBottom: 10 }}
        onChange={(e) => setPassword(e.target.value)}
      />

      <Button type="primary" block onClick={handleLogin}>
        Entrar
      </Button>
    </Card>
  );
}