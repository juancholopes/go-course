package test

// Las funciones pueden recibir mas de dos parámetros y devolver mas de dos resultados

const benefitForWork = 0.04
const taxesCalculated = 0.03

func CalculateTaxes(price float64) float64 {

	return price * 0.12
}

func CalculateSalary(salary float64) (SalaryAfterTaxes float64, SalaryBeforeTaxes float64) {
	return salary + (salary * benefitForWork), salary + (salary * taxesCalculated)
}
