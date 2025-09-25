interface Props {
  text: string,
  img?: string,
  imgClass?: string
}

const Button = ({ text, img, imgClass }: Props) => {
  return (
    <button className="btn-primary py-2 px-4 text-sm font-inter mt-4 flex-center gap-2">
      <img src={img} className={imgClass} />
      {text}
    </button>
  )
}

export default Button