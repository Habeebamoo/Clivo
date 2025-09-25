import { createSlice } from "@reduxjs/toolkit";
import logo from "../../assets/logo.jpg"

export interface Article {
  articleId: string,
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

const initialState = {
  articles: [
    {articleId:"jfif", userPicture: logo, userFullname: "Clivo", username: "@clivo", userVerified: true, title:"How to get a verified account", createdAt: "2 months ago", picture: logo, tags: ["Tech", "Design", "Business"], readTime: "1 mins read time"},
    {articleId: "weio", userPicture: "", userFullname: "Habeeb Amoo", username: "@habeeb_amoo_534", userVerified: false, title:"Go or Rust for backend developement", createdAt: "4 weeks ago", picture: "", tags: ["Tech", "Software"], readTime: "6 mins read time"},
  ]
}

const articleSlice = createSlice({
  name: "articles",
  initialState,
  reducers: {
    setArticles: (state, action) => {

    },
    addArticles: (state, action) => {

    }
  }
})

export const { setArticles } = articleSlice.actions;
export default articleSlice.reducer