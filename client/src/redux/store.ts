import { configureStore } from "@reduxjs/toolkit";
import articleReducer from "./reducers/article_reducer"
import userReducer from "./reducers/user_reducer"

const store = configureStore({
  reducer: {
    articles: articleReducer,
    user: userReducer
  }
})

export default store