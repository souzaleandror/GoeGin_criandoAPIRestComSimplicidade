package databse

var(
	DB *gorm.DB
	err error
)

func ContactComBancoDedados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 ssl=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")	
	}
	DB.AutoMigrate(&models.Aluno{})
}