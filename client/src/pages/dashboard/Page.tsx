import ArticleDisplay from "../../components/ArticleDisplay"
import logo from "../../assets/logo.jpg"
import type { Article } from "../../types/article"

const Dashboard = () => {

  const articles: Article[] = [
    {userPicture: logo, userFullname: "Clivo", username: "@clivo", userVerified: true, title:"How to get a verified account", createdAt: "2 months ago", picture: logo, tags: ["Tech", "Design", "Business"], readTime: "1 mins read time"},
    {userPicture: "", userFullname: "Habeeb Amoo", username: "@habeeb_amoo_534", userVerified: false, title:"Go or Rust for backend developement", createdAt: "4 weeks ago", picture: "", tags: ["Tech", "Software"], readTime: "6 mins read time"},
  ]

  return (
    <main>
      <div className="mt-20"></div>
      
      {/* Posts */}
      <section className="w-[90%] sm:w-[400px] md:w-[500px] mx-auto">
        {articles.map((article: any, i: any) => {
          return (
            <ArticleDisplay key={i} article={article} />
          )
        })}
      </section>
    </main>
  )
}

export default Dashboard