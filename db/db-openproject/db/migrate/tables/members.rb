#-- copyright
# OpenProject is an open source project management software.
# Copyright (C) 2012-2024 the OpenProject GmbH
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License version 3.
#
# OpenProject is a fork of ChiliProject, which is a fork of Redmine. The copyright follows:
# Copyright (C) 2006-2013 Jean-Philippe Lang
# Copyright (C) 2010-2013 the ChiliProject Team
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation; either version 2
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
#
# See COPYRIGHT and LICENSE files for more details.
#++

require_relative 'base'

class Tables::Members < Tables::Base
  def self.table(migration)
    create_table migration do |t|
      t.integer :user_id, default: 0, null: false
      t.integer :project_id, default: 0, null: false
      t.datetime :created_on
      t.boolean :mail_notification, default: false, null: false

      t.index :project_id, name: 'index_members_on_project_id'
      t.index %i[user_id project_id], name: 'index_members_on_user_id_and_project_id', unique: true
      t.index :user_id, name: 'index_members_on_user_id'
    end
  end
end
