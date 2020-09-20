package cmd

import (
	"github.com/chrisbdaemon/contingency/backup"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Perform backup",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		log.Debug("Starting backup")

		b := backup.Backup{}
		b.Size, err = backup.ParseSize(viper.GetString("size"))
		if err != nil {
			log.WithError(err).Fatal("unable to parse size")
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.PersistentFlags().BoolP("pretend", "t", false, "Pretend to backup, print files that will be backed up")
	backupCmd.PersistentFlags().StringP("path", "p", "", "Path to backup")
	backupCmd.PersistentFlags().StringP("size", "s", "4.7GB", "Backup file will not exceed this size")
	viper.BindPFlag("pretend", backupCmd.PersistentFlags().Lookup("pretend"))
	viper.BindPFlag("path", backupCmd.PersistentFlags().Lookup("path"))
}
