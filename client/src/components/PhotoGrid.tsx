import nature from "../assets/nature.jpg"
import commercial from "../assets/commercial.jpg"
import literature from "../assets/literature.jpg"
import tech from "../assets/tech.jpg"
import telephone from "../assets/telephone.jpg"
import chefs from "../assets/chefs.jpg"

const PhotoGrid = () => {
  return (
    <div className="w-[90%] sm:w-[400px] mx-auto">
      <div className="flex-around min-h-[550px]">
        <div className="grid grid-cols-1 gap-1">
          <img src={nature} className="" />
          <img src={commercial} className="" />
        </div>
        <div className="grid grid-cols-1 gap-1 mx-1">
          <img src={literature} className="" />
          <img src={tech} className="" />
        </div>
        <div className="grid grid-cols-1 gap-1">
          <img src={telephone} className="" />
          <img src={chefs} className="" />
        </div>
      </div>
    </div>
  )
}

export default PhotoGrid