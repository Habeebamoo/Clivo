import { useSelector } from "react-redux"
import ArticleDisplay from "../../components/ArticleDisplay"
import type { Article } from "../../redux/reducers/article_reducer"

const Dashboard = () => {
  const articles: Article[] = useSelector((state: any) => state.articles.articles)

  return (
    <main>
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