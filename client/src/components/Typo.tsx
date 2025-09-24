const H1 = (
  { font, color="black", text, others }: 
  { font: string, color?: string, text: string, others: string }
) => {
  return (
    <h1 className={`text-3xl font-${font} text-${color} ${others}`}>{text}</h1>
  )
}

const H2 = (
  { font, color="black", text }: { font: string, color?: string, text: string }
) => {
  return (
    <p className={`text-2xl font-${font} text-${color}`}>{text}</p>
  )
}

export { H1, H2 };