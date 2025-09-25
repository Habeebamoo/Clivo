import type { Article } from "../types/article"
import { H3 } from "./Typo"
import { shorten } from "../utils/utils"
import { IoCheckmarkCircle } from "react-icons/io5"
import { BsShare } from "react-icons/bs"

const ArticleDisplay = ({ article }: { article: Article }) => {
  const toUserProfile = () => {

  }

  const toPost = () => {

  }

  const sharePost = () => {
    
  }

  return (
    <div className="py-4 border-b-1 border-b-muted">
      {/* profile */}
      <div className="hover:bg-mutedLight cursor-pointer p-1">
        <div className="flex-start gap-2">
          {article.userPicture ?
            <div className="h-6 w-6 rounded-full overflow-hidden">
              <img src={article.userPicture} className="h-full w-full object-cover" />
            </div> :
            <div className="w-5 h-5 rounded-full bg-muted border-1 border-accentLight"></div>
          }
          <div className="flex-start gap-1">
            <p className="font-exo text-sm">{article.userFullname}</p>
            {article.userVerified && <IoCheckmarkCircle color="rgba(155, 166, 214, 1)" />}
          </div>
        </div>
        <p className="text-[12px] text-accent pl-7">{article.createdAt}</p>
      </div>

      {/* title & picture */}
      <div className="py-3 my-2 grid grid-cols-6 gap-3 px-1 hover:bg-mutedLight cursor-pointer">
        <div className="col-span-5 break-words">
          <H3 font="inter" text={shorten(article.title, 50)} />
        </div>
        <div className="h-full w-full">
          {article.picture ? (
            <img src={article.picture} className="h-full w-full object-cover" />
          ): (
            <div className="w-full h-full bg-muted object-cover"></div>
          )}
        </div>
      </div>

      {/* tags */}
      <div className="flex-start flex-wrap gap-1 w-[80%]">
        {article.tags?.map((tag: string, i: any) => {
          return <div key={i} className="text-[12px] font-exo py-1 px-2 border-1 border-accent rounded-sm">{tag}</div>
        })}
      </div>

      {/* others */}
      <div className="mt-3 flex-between">
        <p className="text-[12px] font-exo text-accentLight">{article.readTime}</p>
        <div className="text-accentLight">
          <BsShare size={14} />
        </div>
      </div>

    </div>
  )
}

export default ArticleDisplay