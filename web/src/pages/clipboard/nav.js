import {useEffect, useState} from 'react';
import { Menu } from 'antd';
import { useMatches, useNavigate } from 'react-router-dom';
import { get } from 'lodash';
import { ScissorOutlined, SettingOutlined, QuestionCircleOutlined } from '@ant-design/icons';

const NavMenu = (props) => {
    const { sider } = props;
    const [current, setCurrent] = useState('');
    const matches = useMatches();
    const navigate = useNavigate();
    useEffect(( )=> {
        const match = get(matches, matches.length - 1);
        const pathname = get(match, 'pathname');
        setCurrent(pathname);
    }, [matches])

    const items = [
        {
            label: '剪贴板',
            key: '/clip/board',
            icon: <ScissorOutlined />,
        },
        // {
        //     label: '文件对传',
        //     key: 'file',
        //     icon: <PaperClipOutlined />,
        // },
        {
            label: '设置',
            key: '/clip/settings',
            icon: <SettingOutlined />,
        },
        {
            label: <a href="https://github.com/LanceLRQ/cloud-clipboard" target="_blank" rel="noopener noreferrer">关于</a>,
            icon: <QuestionCircleOutlined />,
            key: 'about',
        },
    ];

    const onClick = (e) => {
        if (e.key === 'about') return;
        navigate(e.key);
    };
    return <Menu onClick={onClick} selectedKeys={[current]} mode={sider ? 'vertical' : 'horizontal'} items={items} />;
}

NavMenu.defaultProps = {
    sider: false
}

export default NavMenu;