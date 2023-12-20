import React, { useState } from 'react';
import {Layout, Input, Flex, Button, Space, Avatar, List, Modal } from "antd";
import {
  LikeOutlined, MessageOutlined, StarOutlined,
  SendOutlined, ClearOutlined, CloudUploadOutlined,
} from '@ant-design/icons';

const IconText = ({ icon, text }) => (
  <Space>
    {React.createElement(icon)}
    {text}
  </Space>
);
const data = Array.from({
  length: 23,
}).map((_, i) => ({
  href: 'https://ant.design',
  title: `ant design part ${i}`,
  avatar: `https://xsgames.co/randomusers/avatar.php?g=pixel&key=${i}`,
  description:
    'Ant Design, a design language for background applications, is refined by Ant UED Team.',
  content:
    'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
}));

const PageClipboard = () => {
  const [isInputFocus, setIsInputFocus] = useState(false);
  const [inputText, setInputText] = useState('');
  const [modal] = Modal.useModal();
  const clearText = () => {
    modal.confirm({
      title: 'Confirm',
      content: '确认要清除输入框吗？',
      okText: '确认',
      cancelText: '取消',
      onOk: () => {
        setInputText('');
      }
    });
  }
  return <Layout className="body-layout">
    <Flex justify="space-between" className="header-input-layout">
      {!isInputFocus ? <Flex align="center"><Button icon={<CloudUploadOutlined />}></Button></Flex> : null}
      <Input.TextArea
        onFocus={() => setIsInputFocus(true)}
        onBlur={() => setIsInputFocus(false)}
        autoSize={{ minRows: isInputFocus ? 5 : 1, maxRows: 10 }}
        size="large"
        placeholder="在此输入需要粘贴的内容，或者按Ctrl+V"
        bordered={false}
        value={inputText}
        onChange={(e) => setInputText(e.target.value)}
      />
      <Flex align="end" justify="flex-end" gap={8} vertical={isInputFocus}>
        {isInputFocus ? <Button danger icon={<ClearOutlined />}></Button> : null}
        <Button onClick={clearText} icon={<SendOutlined />}></Button>
      </Flex>
    </Flex>
    <Layout.Content className="content-layout">
      <List
        itemLayout="vertical"
        size="large"
        dataSource={data}
        footer={
          <div>
            <b>ant design</b> footer part
          </div>
        }
        renderItem={(item) => (
          <List.Item
            key={item.title}
            actions={[
              <IconText icon={StarOutlined} text="156" key="list-vertical-star-o" />,
              <IconText icon={LikeOutlined} text="156" key="list-vertical-like-o" />,
              <IconText icon={MessageOutlined} text="2" key="list-vertical-message" />,
            ]}
            extra={
              <img
                width={272}
                alt="logo"
                src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
              />
            }
          >
            <List.Item.Meta
              avatar={<Avatar src={item.avatar} />}
              title={<a href={item.href}>{item.title}</a>}
              description={item.description}
            />
            {item.content}
          </List.Item>
        )}
      />
    </Layout.Content>
  </Layout>
}

export default PageClipboard;