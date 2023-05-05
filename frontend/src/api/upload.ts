import { post } from "./apiClient";

export async function startUploadFlow(params: {
  title: string;
  description: string;
  imageKeywords: string[];
  voice?: string;
  style?: string;
  background?: string;
  textContent: string;
}) {
  const body = {
    ref: "master",
    inputs: {
      title: params.title,
      description: params.description,
      image_keywords: params.imageKeywords.join(","),
      text_content: params.textContent,
      custom_voice: params.voice,
      custom_style: params.style,
      custom_background: params.background,
    },
  };
  return post<string>(`generate`, JSON.stringify(body));
}
