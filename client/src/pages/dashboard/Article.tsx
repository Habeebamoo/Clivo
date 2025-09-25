import { useSelector } from "react-redux"
import { H2 } from "../../components/Typo"
import { useNavigate } from "react-router";
import { useEffect } from "react";
import { articles, type Article } from "../../redux/reducers/article_reducer";
import { shorten } from "../../utils/utils";
import { IoCheckmarkCircle } from "react-icons/io5";
import { FaUpload } from "react-icons/fa";

const Article = () => {
  const articleId: string = useSelector((state: any) => state.articles.activeArticle);
  
  const navigate = useNavigate();

  useEffect(() => {
    if (!articleId) {
      navigate("/dashboard")
    }
  })

  const article: Article | undefined = articles.find((art) => art?.articleId === articleId)

  return (
    <main className="w-[90%] sm:w-[400px] md:w-[500px] mx-auto">
      <div className="mt-20"></div>

      {/* Heading section */}
      <div className="p-1 w-full">
        <div>
          {article!.picture ? (
            <img src={article!.picture} className="h-35 mb-4 mx-auto" />
          ) : (
            <div className="bg-muted w-full flex-grow h-35 mb-4"></div>
          )}
        </div>
        <H2 font="exo" text={shorten(article!.title, 50)} others="text-center" />
      </div>

      {/* Author */}
      <div className="text-[12px] text-accentLight flex-start gap-3 mt-4">
        <p>{article!.readTime}</p>
        <div className="bg-accentLight h-1 w-1 rounded-full"></div>
        <p>{article!.createdAt}</p>
      </div>
      <div className="flex-start gap-3 mt-2">
        <div className="h-7 w-7 rounded-full overflow-hidden">
          {article!.authorPicture ? (
            <img src={article!.authorPicture} className="h-full w-full object-cover" />
          ) : (
            <div className="w-5 h-5 rounded-full bg-muted border-1 border-accentLight"></div>
          )}
        </div>
        <div className="flex-start gap-1">
          <p>{article!.authorFullname}</p>
          {article!.authorVerified && 
            <IoCheckmarkCircle color="rgba(155, 166, 214, 1)" />
          }
        </div>
        <div className="flex-start gap-1 ml-4">
          <button className="btn-primary text-[12px] rounded-full font-inter">Follow</button>

          <button className="text-[12px] flex-center gap-1 border-1 border-accent py-1 px-3 rounded-full cursor-pointer text-black hover:bg-muted ">
          <FaUpload />
          <span>Share</span>
        </button>
        </div>
      </div>
      <hr className="text-muted mt-6 mb-4" />

      {/* Article Content */}
      <div className="font-exo text-text">
        {article!.content}
      </div>

      {/* tags */}
      <div className="mt-10 flex-start gap-2 flex-wrap">
        {article!.tags?.map((tag: string, i: any) => {
          return <div key={i} className="text-[12px] font-exo py-1 px-2 border-1 border-accent rounded-sm">{tag}</div>
        })}
      </div>

    </main>
  )
}

export default Article