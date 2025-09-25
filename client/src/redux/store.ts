import { configureStore } from "@reduxjs/toolkit";
import articleReducer from "./reducers/article_reducer"

const store = configureStore({
  reducer: {
    articles: articleReducer
  }
})

export default store