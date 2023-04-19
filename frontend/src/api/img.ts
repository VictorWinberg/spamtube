import { get } from "./apiClient";

export type ImgData = {
  url: string;
};

export async function getImg() {
  return get<ImgData[]>(`img`);
}
