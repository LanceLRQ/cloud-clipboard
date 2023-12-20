import '@/styles/clipboard.scss';
import {Layout} from 'antd';
import NavMenu from "./nav";
import { Outlet } from 'react-router-dom';

function Index() {
  return (
    <Layout className="main-layout">
      <Layout.Sider
        breakpoint="lg"
        collapsedWidth="0"
        className="sider-layout"
      >
        <NavMenu sider/>
      </Layout.Sider>
      <Outlet></Outlet>
    </Layout>
  );
}

export default Index;
