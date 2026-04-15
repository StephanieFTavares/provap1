import { Card, Button, List, Typography } from "antd";
import { useEffect, useState } from "react";
import api from "../services/api";

const { Title } = Typography;

export default function Admin() {
  const [users, setUsers] = useState([]);

  const loadUsers = async () => {
    try {
      const res = await api.get("/users", {
        headers: { Authorization: "token" },
      });
      setUsers(res.data);
    } catch {
      console.log("Erro ao carregar");
    }
  };

  const approve = async (id) => {
    await api.put(`/approve/${id}`, {}, {
      headers: { Authorization: "token" },
    });
    loadUsers();
  };

  useEffect(() => {
    loadUsers();
  }, []);

  return (
    <Card style={{ maxWidth: 600, margin: "20px auto" }}>
      <Title level={3}>Painel Admin</Title>

      <List
        dataSource={users}
        renderItem={(user) => (
          <List.Item
            actions={[
              !user.active && (
                <Button onClick={() => approve(user.id)}>
                  Aprovar
                </Button>
              ),
            ]}
          >
            {user.name} - {user.active ? "Ativo" : "Pendente"}
          </List.Item>
        )}
      />
    </Card>
  );
}