const timeStringToBrDateString = (strDate: string) => {
  const dt = new Date(strDate)
  return `${dt.toLocaleDateString('pt-BR')} ${dt.toLocaleTimeString('pt-BR')}`
}

const dateToBrDateString = (date: Date) => {
  return `${date.toLocaleDateString('pt-BR')} ${date.toLocaleTimeString('pt-BR')}`
}

export { timeStringToBrDateString, dateToBrDateString }
