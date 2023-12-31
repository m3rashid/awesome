import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

export const baseUrl = 'http://127.0.0.1:4000';
// export const baseUrl = 'http://192.168.141.113:4000';

export type RequestOptions = Omit<AxiosRequestConfig, 'url' | 'baseURL'>;
export type OtherOptions = {
  token?: string;
  noBaseURL?: boolean;
};

export const service = <Res = any>(
  url: string,
  options: RequestOptions,
  OtherOptions: OtherOptions = {}
) => {
  return (config?: AxiosRequestConfig): Promise<AxiosResponse<Res>> => {
    return axios<Res>({
      url,
      method: options.method || 'GET',
      baseURL: OtherOptions.noBaseURL ? '' : baseUrl,
      ...(config || {}),
      headers: {
        ...(options.headers || {}),
        ...(OtherOptions.token
          ? { Authorization: `Bearer ${OtherOptions.token}` }
          : {}),
      },
    });
  };
};

export type Service<Res> = typeof service<Res>;
