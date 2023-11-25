import { Card, Typography } from 'antd';
import React from 'react';

import { appPaths } from '../../constants/paths';

type AppSwitchDrawerContentsProps = {
  onRouteClick: (route: string) => void;
};

const AppSwitchDrawerContents: React.FC<AppSwitchDrawerContentsProps> = ({
  onRouteClick,
}) => {
  return (
    <>
      <Typography.Title level={3} type='secondary' style={{ marginLeft: 8 }}>
        Apps and Services
      </Typography.Title>

      {appPaths.map((routeConfig) => {
        return (
          <Card
            key={routeConfig.route}
            style={{ marginBottom: '1rem' }}
            bodyStyle={{
              padding: 8,
              paddingLeft: 16,
              paddingRight: 16,
              cursor: 'pointer',
            }}
          >
            <div
              className='flex items-center gap-2'
              onClick={() => onRouteClick(routeConfig.route)}
            >
              <routeConfig.icon
                style={{ fontSize: 16, margin: 0, padding: 0 }}
              />

              <Typography.Title level={4}>{routeConfig.name}</Typography.Title>
            </div>
            <Typography.Text>{routeConfig.description}</Typography.Text>

            <div className='flex flex-wrap'>
              {routeConfig.children &&
                routeConfig.children.map((child) => {
                  return (
                    <Typography.Link
                      strong
                      key={child.route}
                      className='mr-5'
                      onClick={() => onRouteClick(child.route)}
                    >
                      {child.name}
                    </Typography.Link>
                  );
                })}
            </div>
          </Card>
        );
      })}
    </>
  );
};

export default AppSwitchDrawerContents;
