import { Spinner } from '@fluentui/react-components';
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import PostCard from '../../../components/atoms/postCard';
import PageContainer from '../../../components/pageContainer';
import { service } from '../../../helpers/service';

const PostDetails: React.FC = () => {
  const params = useParams();
  const [post, setPost] = useState<any | null>(null);

  const getPostDetailsService = service('/api/community/post/get', {
    method: 'POST',
  });

  const getPostDetails = async (postId: string) => {
    const response = await getPostDetailsService({
      data: {
        searchCriteria: { id: Number(postId) },
        populate: ['Topic', 'User'],
      },
    });
    setPost(response.data);
  };

  useEffect(() => {
    if (!params.postId) return;
    getPostDetails(params.postId).catch(console.log);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [params.postId]);

  if (!post) return <Spinner size='extra-large' />;

  return (
    <PageContainer>
      <div className='flex items-center justify-center mt-2'>
        <PostCard post={post} type='detail' />
      </div>
    </PageContainer>
  );
};

export default PostDetails;
