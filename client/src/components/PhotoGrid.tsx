import nature from "../assets/nature.jpg"
import commercial from "../assets/commercial.jpg"
import literature from "../assets/literature.jpg"
import tech from "../assets/tech.jpg"
import telephone from "../assets/telephone.jpg"
import chefs from "../assets/chefs.jpg"

const PhotoGrid = () => {
  return (
    <div className="w-[90%] sm:w-[400px] mx-auto">
      <div className="grid grid-cols-3 gap-2">

        <div className="flex flex-col gap-2 justify-center">
          <img src={commercial} className="aspect-square object-cover w-full" />
          <img src={nature} className="aspect-square object-cover w-full" />
        </div>

        <div className="flex flex-col gap-2">
          <img src={chefs} className="aspect-square object-cover w-full" />
          <img src={tech} className="object-cover w-full" style={{ aspectRatio: "1 / 2"}} />
        </div>

        <div className="flex flex-col gap-2 justify-center">
          <img src={telephone} className="aspect-square object-cover w-full" />
          <img src={literature} className="aspect-square object-cover w-full" />
        </div>

      </div>
    </div>
  )
}

export default PhotoGrid