import './app.scss';
import {ConfigProvider, Layout} from 'antd';
import antdConf from "./antd_conf";
import NavMenu from "./nav";

function App() {
  return (
    <ConfigProvider {...antdConf}>
      <Layout className="main-layout">
        <Layout.Sider
          breakpoint="lg"
          collapsedWidth="0"
          className="sider-layout"
        >
          <NavMenu sider/>
        </Layout.Sider>
        <Layout className="body-layout">
          <Layout.Header className="header-layout"></Layout.Header>
          <Layout.Content className="content-layout">

          </Layout.Content>
        </Layout>
      </Layout>
    </ConfigProvider>
  );
}

export default App;
