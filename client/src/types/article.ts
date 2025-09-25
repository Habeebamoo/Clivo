export interface Article {
  userPicture: string,
  userFullname: string,
  username: string,
  userVerified: boolean,
  createdAt: string,
  title: string,
  picture?: string,
  tags?: string[],
  readTime: string
}