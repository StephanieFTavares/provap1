import { Layout, Typography } from "antd";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Admin from "./pages/Admin";

const { Header, Content } = Layout;
const { Title } = Typography;

function App() {
  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Header style={{ background: "#001529" }}>
        <Title style={{ color: "#fff", margin: 0 }} level={3}>
          Sistema Condlink
        </Title>
      </Header>

      <Content style={{ padding: "40px" }}>
        <Login />
        <Register />
        <Admin />
      </Content>
    </Layout>
  );
}

export default App;