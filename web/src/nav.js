import { useState } from 'react';
import { Menu } from 'antd';
import { ScissorOutlined, SettingOutlined, PaperClipOutlined, QuestionCircleOutlined } from '@ant-design/icons';


const NavMenu = (props) => {
    const { sider } = props;
    const [current, setCurrent] = useState('mail');
    const items = [
        {
            label: '剪贴板',
            key: 'clipboard',
            icon: <ScissorOutlined />,
        },
        {
            label: '文件对传',
            key: 'file',
            icon: <PaperClipOutlined />,
        },
        {
            label: '设置',
            key: 'settings',
            icon: <SettingOutlined />,
        },
        {
            label: <a href="https://github.com/LanceLRQ/cloud-clipboard" target="_blank" rel="noopener noreferrer">关于</a>,
            icon: <QuestionCircleOutlined />,
            key: 'about',
        },
    ];

    const onClick = (e) => {
        setCurrent(e.key);
    };
    return <Menu onClick={onClick} selectedKeys={[current]} mode={sider ? 'vertical' : 'horizontal'} items={items} />;
}

NavMenu.defaultProps = {
    sider: false
}

export default NavMenu;