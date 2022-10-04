declare global {
  declare module "*.svg" {
    const filePath: string;
    export default filePath;
  }

  declare module "*.png" {
    const value: string;
    export = value;
  }
}

export {};
