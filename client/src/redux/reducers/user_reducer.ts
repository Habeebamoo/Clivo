import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  user: {
    name: "Habeeb Amoo",
    email: "",
    picture: "",
    bio: "Clivo CEO",
    username: "@habeeb_amoo_414",
    following: 16,
    followers: 400,
    verified: true,
    profileLink: "habeebamoo.netlify.app",
    createdAt: "3 months ago"
  }
}

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (state: any, action: any) => {

    }
  }
})

export const { setUser } = userSlice.actions;
export default userSlice.reducer;