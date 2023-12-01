import { EditOutlined } from '@ant-design/icons';
import { useAuthValue } from '@awesome/shared/atoms/auth';
import {
  Button,
  Card,
  Form,
  Input,
  message,
  Modal,
  Select,
  Typography,
} from 'antd';
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import Loader from '../../../components/atoms/loader';
import PageContainer from '../../../components/pageContainer';
import { service } from '../../../helpers/service';

const PostDetails: React.FC = () => {
  const auth = useAuthValue();
  const params = useParams();
  const [form] = Form.useForm();
  const [topics, setTopics] = useState<any | null>(null);
  const [modalOpen, setModalOpen] = useState(false);
  const [post, setPost] = useState<any | null>(null);

  const getTopicsService = service('/api/community/topics');
  const updatePostService = service('/api/community/post/update');
  const getPostDetailsService = service('/api/community/post/get');

  const getPostDetails = async (postId: string) => {
    const response = await getPostDetailsService({
      method: 'POST',
      data: { id: Number(postId) },
    });
    setPost(response.data);
  };

  const getTopics = async () => {
    const response = await getTopicsService({
      method: 'POST',
      data: {
        searchCriteria: { deleted: false },
        paginationOptions: { limit: 10, page: 1 },
      },
    });
    setTopics(response.data);
  };

  const handleEditPost = async () => {
    try {
      await form.validateFields();
      await updatePostService({
        method: 'POST',

        data: {
          searchCriteria: { id: post.id },
          update: {
            ...form.getFieldsValue(),
          },
        },
      });
      form.resetFields();
      setModalOpen(false);
      if (params.postId) getPostDetails(params.postId).catch(console.log);
    } catch (err: any) {
      console.log(err);
      message.error('Error in updating post');
    }
  };

  useEffect(() => {
    if (!params.postId) return;
    Promise.allSettled([getTopics(), getPostDetails(params.postId)]).catch(
      console.log
    );
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [params.postId]);

  if (!post) return <Loader />;

  return (
    <PageContainer>
      <Modal
        footer={null}
        open={modalOpen}
        title='Edit new topic'
        onCancel={() => setModalOpen(false)}
      >
        <Form
          form={form}
          layout='vertical'
          initialValues={post}
          onFinish={handleEditPost}
        >
          <Form.Item name='title' label='Title'>
            <Input placeholder='Enter title for your post' />
          </Form.Item>

          <Form.Item name='body' label='Body'>
            <Input.TextArea placeholder='Your article here' />
          </Form.Item>

          <Form.Item name='topicId' label='Topic'>
            <Select placeholder='Select suitable topic for your article'>
              {(topics?.docs || []).map((topic: any) => {
                return (
                  <Select.Option key={topic.id} value={topic.id}>
                    {topic.name}
                  </Select.Option>
                );
              })}
            </Select>
          </Form.Item>

          <div className='flex gap-2 justify-end'>
            <Button type='primary' htmlType='submit' icon={<EditOutlined />}>
              Update Post
            </Button>
            <Button onClick={() => setModalOpen(false)}>Cancel</Button>
          </div>
        </Form>
      </Modal>

      <div className='flex items-center justify-center mt-2'>
        <Card
          {...{
            style: { minWidth: 320, maxWidth: 640 },
            ...(post.title ? { title: post.title } : {}),
            ...(post.userId === auth?.user.id
              ? {
                  extra: (
                    <Button
                      type='link'
                      icon={<EditOutlined />}
                      onClick={() => setModalOpen(true)}
                    />
                  ),
                }
              : {}),
          }}
        >
          <Typography.Text>{post.body}</Typography.Text>
        </Card>
      </div>
    </PageContainer>
  );
};

export default PostDetails;
