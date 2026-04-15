import { Card, Input, Button, Typography } from "antd";
import { useState } from "react";
import api from "../services/api";

const { Title } = Typography;

export default function Register() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleRegister = async () => {
    try {
      await api.post("/register", { name, email, password });
      alert("Cadastro enviado!");
    } catch {
      alert("Erro ao cadastrar");
    }
  };

  return (
    <Card style={{ maxWidth: 400, margin: "20px auto" }}>
      <Title level={3}>Cadastro</Title>

      <Input
        placeholder="Nome"
        style={{ marginBottom: 10 }}
        onChange={(e) => setName(e.target.value)}
      />

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

      <Button type="primary" block onClick={handleRegister}>
        Cadastrar
      </Button>
    </Card>
  );
}